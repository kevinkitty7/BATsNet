#pragma once

#include <errno.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/epoll.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <unistd.h>

#include <array>
#include <chrono>
#include <mutex>
#include <queue>
#include <thread>
#include <unordered_map>
#include <unordered_set>

#include "include/common.h"
#include "proto_gen/scheduler.pb.h"

namespace scheduler {
constexpr int kMissionWorker = 2;

/****
 * Two type of mission:
 * 1. Real Time Mission
 * 2. Batch Mssion
 *
 * The (1) mission should start asap
 * Don't need to consider the resource
 * or other contraint
 *
 * The (2) mission need to be consider the resource
 * it need to wait until resource available.
 * When multiple (2) mission need to schedule,
 * schedule to maximize thoughput
 */

struct ResourceTable {
  int total_core;
  int avail_core;

  int total_mem;
  int avail_mem;

  int total_gpu;
  int avail_gpu;
};

struct ScheduledMission {
  ScheduledMission() {}
  ScheduledMission(const ScheduledMission &other) { *this = other; }
  ScheduledMission &operator=(const ScheduledMission &other) {
    name = other.name;
    priority = other.priority;
    start_s_in_day = other.start_s_in_day;
    real_time = other.real_time;
    setting.CopyFrom(other.setting);
    return *this;
  }
  std::string name;
  int priority = 0;
  int start_s_in_day = 0;
  bool real_time = false;
  MissionSetting setting;
};

struct PendingCompare {
  bool operator()(const ScheduledMission &m1, const ScheduledMission &m2) {
    return m1.start_s_in_day < m2.start_s_in_day;
  }
};

struct WaitingCompare {
  bool operator()(const ScheduledMission &m1, const ScheduledMission &m2) {
    return m1.priority > m2.priority;
  }
};

class Scheduler {
public:
  Scheduler(const SchedulerSetting &setting);
  ~Scheduler();
  void Setup();
  void Run();

private:
  // Interactive to Mission
  void AcceptConnections();
  void ListenToMissions();
  void MissionLoop();

  void HandlePingPacket(PingPacket *packet, ClientConnection *conn);
  void HandleRegisterPacket(RegisterPacket *packet, ClientConnection *conn);
  void HandleAllowStartAck(AllowStartAck *packet, ClientConnection *conn);
  void HandleMissionDone(MissionDone *packet, ClientConnection *conn);

  // Maintain after interact to Mission
  bool IsStillValid(const ClientConnection *conn) const;
  void RemoveFromScheduling(const ClientConnection *);

  // The scheduling
  void TimeLoop();
  bool AbleToSelect(const ScheduledMission &mission) const;
  void TrySelectWaitingMission();
  void AddRunningMission(const ScheduledMission &mission);
  void ReleaseMissionResource(const ScheduledMission &mission);

  bool running_ = true;
  int server_sock_;
  SchedulerSetting setting_;

  static constexpr int kMaxEvents = 1024;
  int epoll_fd_;

  std::thread accept_thread_;
  std::thread mission_thread_;

  std::array<std::thread, kMissionWorker> mission_workers_;
  std::array<std::queue<std::shared_ptr<ClientConnection>>, kMissionWorker>
      mission_socks_;

  std::mutex need_handle_mtx_;
  std::queue<std::shared_ptr<ClientConnection>> need_handle_sock_;

  std::mutex all_conns_mtx_;
  std::unordered_map<int, std::shared_ptr<ClientConnection>> all_conns_;
  // The register conn
  std::mutex conn_map_mtx_;
  std::unordered_map<std::string, std::shared_ptr<ClientConnection>> conn_map_;
  std::array<std::mutex, kMissionWorker> mtxs_;
  PacketHandleMap handle_map_;

  std::thread schedule_thread_;
  ResourceTable resource_tbl_;
  std::mutex mission_lock_;

  // The mission that should not start yet
  std::priority_queue<ScheduledMission, std::vector<ScheduledMission>,
                      PendingCompare>
      pending_mission_;
  // The mission that should start, but not gain the start permission yet
  std::priority_queue<ScheduledMission, std::vector<ScheduledMission>,
                      WaitingCompare>
      waiting_mission_;
  // The started mission
  std::unordered_map<std::string, ScheduledMission> running_mission_;
};

} // namespace scheduler
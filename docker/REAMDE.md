# To enable build this arm image in x86
sudo apt-get install qemu binfmt-support qemu-user-static

# Cmd to build
docker build -t batsnet -f docker/Dockerfile ./

# Cmd to Submit task
kubectl apply -f demo_image_detection.yaml
echo "EXPERIMENT 1"
./cleaningDocker.sh

sleep 30

./concurrentDocker.sh

sleep 60

echo "EXPERIMENT 2"
./cleaningDocker.sh

sleep 30

./concurrentDocker.sh

sleep 60

echo "EXPERIMENT 3"
./cleaningDocker.sh

sleep 30

./concurrentDocker.sh

sleep 60

echo "EXPERIMENT 4"
./cleaningDocker.sh

sleep 30

./concurrentDocker.sh

sleep 60

echo "EXPERIMENT 5"
./cleaningDocker.sh

sleep 30

./concurrentDocker.sh

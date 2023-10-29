echo "EXPERIMENT 1"
./cleaningDistroMash.sh

sleep 30

./concurrentDistroMash2.sh

sleep 200

echo "EXPERIMENT 2"
./cleaningDistroMash.sh

sleep 30

./concurrentDistroMash2.sh

sleep 200

echo "EXPERIMENT 3"
./cleaningDistroMash.sh

sleep 30

./concurrentDistroMash2.sh

sleep 200

echo "EXPERIMENT 4"
./cleaningDistroMash.sh

sleep 30

./concurrentDistroMash2.sh

sleep 200

echo "EXPERIMENT 5"
./cleaningDistroMash.sh

sleep 30

./concurrentDistroMash2.sh

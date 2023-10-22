(ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "time docker pull 10.156.0.19:5009/700mbv1:latest > /dev/null 2>&1" && echo "vm-4") &
(ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "time docker pull 10.156.0.19:5009/700mbv1:latest > /dev/null 2>&1" && echo "vm-3") &
(ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "time docker pull 10.156.0.19:5009/700mbv1:latest > /dev/null 2>&1" && echo "vm-2") &

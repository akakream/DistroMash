echo "vm-2:"
(sleep 10 && ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "time sudo crictl pull docker.io/akakream/700mbv1:latest > /dev/null 2>&1") &
echo "vm-3:"
(sleep 180 && ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "time sudo crictl pull docker.io/akakream/700mbv1:latest > /dev/null 2>&1") &
echo "vm-4:"
(sleep 270 && ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "time sudo crictl pull docker.io/akakream/700mbv1:latest > /dev/null 2>&1") &
(sleep 280 && echo "DONE") &

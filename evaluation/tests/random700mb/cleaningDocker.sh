ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "docker rmi 10.156.0.19:5009/700mbv1:latest &>/dev/null"
echo "vm-2 cleaned"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "docker rmi 10.156.0.19:5009/700mbv1:latest &>/dev/null"
echo "vm-3 cleaned"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "docker rmi 10.156.0.19:5009/700mbv1:latest &>/dev/null"
echo "vm-4 cleaned"


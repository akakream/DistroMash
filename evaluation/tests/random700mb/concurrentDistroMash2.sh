#curl -X 'POST' \
#  'http://localhost:3000/api/v1/strategy' \
#  -H 'accept: application/json' \
#  -H 'Content-Type: application/json' \
#  -d '{
#  "execute": true,
#  "nametag": "redis:7.2",
#  "percentage": 100,
#  "target": "",
#  "type": "target"
#}'

curl -X 'POST' \
  'http://localhost:3000/api/v1/strategy' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "execute": true,
  "nametag": "700mbv1:latest",
  "percentage": 100,
  "target": "",
  "type": "percentage"
}'

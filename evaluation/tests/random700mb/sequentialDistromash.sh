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


(sleep 10 && curl -X 'POST' \
  'http://localhost:3000/api/v1/strategy' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "execute": true,
  "nametag": "700mbv1:latest",
  "percentage": 100,
  "target": "QmceJNPxC2VgJWriywA7Lk7wyWquj4tMyZRMJDMYZuXFtT",
  "type": "target"
}') &

(sleep 130 && curl -X 'POST' \
  'http://localhost:3000/api/v1/strategy' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "execute": true,
  "nametag": "700mbv1:latest",
  "percentage": 100,
  "target": "QmXMF4FCbfHs7wHk1HUmjaw2oWZJxZ9W75oJSjX96eKi21",
  "type": "target"
}') &

(sleep 250 && curl -X 'POST' \
  'http://localhost:3000/api/v1/strategy' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "execute": true,
  "nametag": "700mbv1:latest",
  "percentage": 100,
  "target": "QmUa4i4Yt5LMp9ZeuRCvBdGZTJvihSc7f32bazmRVGtZVc",
  "type": "target"
}') &

(sleep 260 && echo "DONE") &

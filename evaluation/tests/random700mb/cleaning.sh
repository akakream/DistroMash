#curl -X 'DELETE' \
#  'http://localhost:3000/api/v1/crdt/QmcpQ1HdTw7QuGpmrkkBvYoDNE16Lgh1w9EyoktWgPQrbw' \
#  -H 'accept: application/json'

curl -X 'DELETE' \
  'http://localhost:3000/api/v1/crdt/QmceJNPxC2VgJWriywA7Lk7wyWquj4tMyZRMJDMYZuXFtT' \
  -H 'accept: application/json'

curl -X 'DELETE' \
  'http://localhost:3000/api/v1/crdt/QmXMF4FCbfHs7wHk1HUmjaw2oWZJxZ9W75oJSjX96eKi21' \
  -H 'accept: application/json'

curl -X 'DELETE' \
  'http://localhost:3000/api/v1/crdt/QmUa4i4Yt5LMp9ZeuRCvBdGZTJvihSc7f32bazmRVGtZVc' \
  -H 'accept: application/json'


#curl -X 'DELETE' \
#  'http://localhost:3000/api/v1/crdt/target-redis%3A7.2-QmcpQ1HdTw7QuGpmrkkBvYoDNE16Lgh1w9EyoktWgPQrbw-active' \
#  -H 'accept: application/json'

curl -X 'DELETE' \
  'http://localhost:3000/api/v1/crdt/target-700mbv1%3Alatest-QmceJNPxC2VgJWriywA7Lk7wyWquj4tMyZRMJDMYZuXFtT-active' \
  -H 'accept: application/json'

curl -X 'DELETE' \
  'http://localhost:3000/api/v1/crdt/target-700mbv1%3Alatest-QmXMF4FCbfHs7wHk1HUmjaw2oWZJxZ9W75oJSjX96eKi21-active' \
  -H 'accept: application/json'

curl -X 'DELETE' \
  'http://localhost:3000/api/v1/crdt/target-700mbv1%3Alatest-QmUa4i4Yt5LMp9ZeuRCvBdGZTJvihSc7f32bazmRVGtZVc-active' \
  -H 'accept: application/json'

# ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "docker rmi redis:7.2"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "docker rmi 700mbv1:latest"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "docker rmi 700mbv1:latest"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "docker rmi 700mbv1:latest"

# ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "ipfs pin rm bafybeidk54vljmbowp7s7thkhydm3zwwom6v5mupeyqsonvpqm6ruxovvi"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs pin rm bafybeieguswvwkzkpz5pl2izv3zeaxp22oxsedmflo5day4r5xeozcx4sq"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs pin rm bafybeieguswvwkzkpz5pl2izv3zeaxp22oxsedmflo5day4r5xeozcx4sq"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs pin rm bafybeieguswvwkzkpz5pl2izv3zeaxp22oxsedmflo5day4r5xeozcx4sq"

# ssh -i ~/.ssh/gcp-instance-1 akakream@34.107.53.120 "ipfs repo gc"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.242.230.77 "ipfs repo gc &>/dev/null"
echo "ipfs repo gc ran for 2nd instance"
ssh -i ~/.ssh/gcp-instance-1 akakream@35.198.66.248 "ipfs repo gc &>/dev/null"
echo "ipfs repo gc ran for 3rd instance"
ssh -i ~/.ssh/gcp-instance-1 akakream@34.159.130.93 "ipfs repo gc &>/dev/null"
echo "ipfs repo gc ran for 4th instance"

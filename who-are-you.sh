#! bin/bash

curl -s  https://learn.01founders.co/assets/superhero/all.json |jq '.[] |select( .id ==70 ).name'


# jq '.[] |select( .id ==70 ).name'

#jq '.[51].name'

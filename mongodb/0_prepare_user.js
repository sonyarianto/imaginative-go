db.auth('root', 'mongodbpassword')

db.getSiblingDB('admin')

db.createUser({user:'root', pwd:'mongodbpassword', roles:[{role:'userAdmin', db:'go_db'}, {role:'readWrite', db:'go_db'}]})
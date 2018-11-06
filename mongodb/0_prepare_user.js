db.auth('root', 'mongodbpassword')

db.getSiblingDB('go_db').createUser({user:'root', pwd:'mongodbpassword', roles:[{role:'userAdminAnyDatabase', db:'go_db'}]})
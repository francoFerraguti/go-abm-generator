import axios from 'axios';

const URL_BASE = 'localhost:8083';

export default {
    createUser(user) {
return axios.post(URL_BASE + '/user', user, {});
},

    editUser(idUser, user) {
return axios.put(URL_BASE + '/user/id' + idUser, user, {});
},

    deleteUser(idUser) {
return axios.delete(URL_BASE + '/user/id' + idUser, {});
},

    getAllUser() {
return axios.get(URL_BASE + '/user/all', {});
},

}
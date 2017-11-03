import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createUser(user) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/user', user, {});
},

    editUser(idUser, user) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/user/id/' + idUser, user, {});
},

    deleteUser(idUser) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/user/id/' + idUser, {});
},

    getAllUser() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/user/list', {});
},

}
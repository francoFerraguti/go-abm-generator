import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createRole(role) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/role', role, {});
},

    editRole(idRole, role) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/role/id' + idRole, role, {});
},

    deleteRole(idRole) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/role/id' + idRole, {});
},

    getAllRole() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/role/list', {});
},

}
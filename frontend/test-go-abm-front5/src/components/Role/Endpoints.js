import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createRole(role) {
return axios.post(URL_BASE + '/role', role, {});
},

    editRole(idRole, role) {
return axios.put(URL_BASE + '/role/id' + idRole, role, {});
},

    deleteRole(idRole) {
return axios.delete(URL_BASE + '/role/id' + idRole, {});
},

    getAllRole() {
return axios.get(URL_BASE + '/role/list', {});
},

}
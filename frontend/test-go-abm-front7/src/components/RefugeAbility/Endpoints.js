import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createRefugeAbility(refugeAbility) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/refugeAbility', refugeAbility, {});
},

    editRefugeAbility(idRefugeAbility, refugeAbility) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/refugeAbility/id' + idRefugeAbility, refugeAbility, {});
},

    deleteRefugeAbility(idRefugeAbility) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/refugeAbility/id' + idRefugeAbility, {});
},

    getAllRefugeAbility() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/refugeAbility/list', {});
},

}
import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createPet(pet) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/pet', pet, {});
},

    editPet(idPet, pet) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/pet/id' + idPet, pet, {});
},

    deletePet(idPet) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/pet/id' + idPet, {});
},

    getAllPet() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/pet/list', {});
},

}
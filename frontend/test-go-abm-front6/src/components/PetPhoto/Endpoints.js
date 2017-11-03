import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createPetPhoto(petPhoto) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/petPhoto', petPhoto, {});
},

    editPetPhoto(idPetPhoto, petPhoto) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/petPhoto/id' + idPetPhoto, petPhoto, {});
},

    deletePetPhoto(idPetPhoto) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/petPhoto/id' + idPetPhoto, {});
},

    getAllPetPhoto() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/petPhoto/list', {});
},

}
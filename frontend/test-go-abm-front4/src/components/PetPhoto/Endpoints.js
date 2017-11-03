import axios from 'axios';

const URL_BASE = 'localhost:8083';

export default {
    createPetPhoto(petPhoto) {
return axios.post(URL_BASE + '/petPhoto', petPhoto, {});
},

    editPetPhoto(idPetPhoto, petPhoto) {
return axios.put(URL_BASE + '/petPhoto/id' + idPetPhoto, petPhoto, {});
},

    deletePetPhoto(idPetPhoto) {
return axios.delete(URL_BASE + '/petPhoto/id' + idPetPhoto, {});
},

    getAllPetPhoto() {
return axios.get(URL_BASE + '/petPhoto/all', {});
},

}
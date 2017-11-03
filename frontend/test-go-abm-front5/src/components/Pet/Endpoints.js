import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createPet(pet) {
return axios.post(URL_BASE + '/pet', pet, {});
},

    editPet(idPet, pet) {
return axios.put(URL_BASE + '/pet/id' + idPet, pet, {});
},

    deletePet(idPet) {
return axios.delete(URL_BASE + '/pet/id' + idPet, {});
},

    getAllPet() {
return axios.get(URL_BASE + '/pet/list', {});
},

}
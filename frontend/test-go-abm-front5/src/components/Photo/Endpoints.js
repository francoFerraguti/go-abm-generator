import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createPhoto(photo) {
return axios.post(URL_BASE + '/photo', photo, {});
},

    editPhoto(idPhoto, photo) {
return axios.put(URL_BASE + '/photo/id' + idPhoto, photo, {});
},

    deletePhoto(idPhoto) {
return axios.delete(URL_BASE + '/photo/id' + idPhoto, {});
},

    getAllPhoto() {
return axios.get(URL_BASE + '/photo/list', {});
},

}
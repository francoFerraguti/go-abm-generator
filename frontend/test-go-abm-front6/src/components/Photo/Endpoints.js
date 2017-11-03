import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createPhoto(photo) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/photo', photo, {});
},

    editPhoto(idPhoto, photo) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/photo/id' + idPhoto, photo, {});
},

    deletePhoto(idPhoto) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/photo/id' + idPhoto, {});
},

    getAllPhoto() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/photo/list', {});
},

}
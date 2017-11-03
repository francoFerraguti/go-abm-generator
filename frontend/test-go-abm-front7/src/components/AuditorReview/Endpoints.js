import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createAuditorReview(auditorReview) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.post(URL_BASE + '/auditorReview', auditorReview, {});
},

    editAuditorReview(idAuditorReview, auditorReview) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.put(URL_BASE + '/auditorReview/id' + idAuditorReview, auditorReview, {});
},

    deleteAuditorReview(idAuditorReview) {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.delete(URL_BASE + '/auditorReview/id' + idAuditorReview, {});
},

    getAllAuditorReview() {
axios.defaults.headers.common['Authorization'] = localStorage.getItem('token');
return axios.get(URL_BASE + '/auditorReview/list', {});
},

}
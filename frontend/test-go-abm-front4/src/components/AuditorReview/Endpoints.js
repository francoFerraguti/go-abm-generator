import axios from 'axios';

const URL_BASE = 'localhost:8083';

export default {
    createAuditorReview(auditorReview) {
return axios.post(URL_BASE + '/auditorReview', auditorReview, {});
},

    editAuditorReview(idAuditorReview, auditorReview) {
return axios.put(URL_BASE + '/auditorReview/id' + idAuditorReview, auditorReview, {});
},

    deleteAuditorReview(idAuditorReview) {
return axios.delete(URL_BASE + '/auditorReview/id' + idAuditorReview, {});
},

    getAllAuditorReview() {
return axios.get(URL_BASE + '/auditorReview/all', {});
},

}
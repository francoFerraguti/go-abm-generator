import axios from 'axios';

const URL_BASE = 'http://localhost:8083';

export default {
    createRefugeAbility(refugeAbility) {
return axios.post(URL_BASE + '/refugeAbility', refugeAbility, {});
},

    editRefugeAbility(idRefugeAbility, refugeAbility) {
return axios.put(URL_BASE + '/refugeAbility/id' + idRefugeAbility, refugeAbility, {});
},

    deleteRefugeAbility(idRefugeAbility) {
return axios.delete(URL_BASE + '/refugeAbility/id' + idRefugeAbility, {});
},

    getAllRefugeAbility() {
return axios.get(URL_BASE + '/refugeAbility/all', {});
},

}
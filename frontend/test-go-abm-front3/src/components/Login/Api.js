import axios from 'axios';

const URL_BASE = 'localhost:8083';

export default {
    loginAccess(user) {
        return axios.post(URL_BASE + '/login', user, {});
    },
};

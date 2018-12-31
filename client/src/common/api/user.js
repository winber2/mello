import axios from 'axios';
import { USER, API, TOKEN } from './constants';

export function getUsers() {
  return axios.get(`/${API}/${USER}/`).then(({ data }) => data);
}

export function postUser(user) {
  return axios.post(`/${API}/${USER}/`, user).then(({ data }) => data);
}

export function getToken(params) {
  return axios.get(`/${TOKEN}`, { params }).then(({ data }) => {
    axios.defaults.headers.common.Authorization = data;
  });
}

import axios from 'axios';
import { USERS, API } from 'common/constants/app';

export function postUser(params) {
  return axios.post(`/${API}/${USERS}/`, params);
}

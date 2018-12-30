import axios from 'axios';
import { USER, API } from 'common/constants/app';

export function postUser(user) {
  return axios.post(`/${API}/${USER}/`, user);
}

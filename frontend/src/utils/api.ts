import { Axios, type AxiosInstance } from 'axios'

export const axios: Axios = new Axios({
	baseURL: 'http://localhost:3000/api',
	withCredentials: true,
})
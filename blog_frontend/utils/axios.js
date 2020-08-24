import axios from "axios";
console.log(process.env.NODE_ENV);
const instance = axios.create({
	baseURL:
		process.env.NODE_ENV !== "development"
			? "http://webblog-env.eba-km92i85a.us-east-2.elasticbeanstalk.com/"
			: "http://localhost:8010/",
	headers: { "Content-Type": "application/json" },
});
instance.interceptors.response.use(
	function (response) {
		return response.data;
	},
	function (error) {
		if (error.response && error.response.status === 401) {
			window.location = `/auth/login?path=${window.location.pathname}`;
			return false;
		}
		return Promise.reject(error);
	}
);
export default instance;

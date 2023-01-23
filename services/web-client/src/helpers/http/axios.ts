import axios from "axios";
import config from "config/vars"

export default axios.create({
    baseURL: config.BASE_URL
})
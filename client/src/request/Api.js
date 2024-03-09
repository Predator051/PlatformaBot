import axios, {AxiosHeaders} from "axios";

export async function SendPost(url, data) {
    return await axios.post(`http://${process.env.REACT_APP_PB_SERVER_IP}:8080/${url}`, {
        ...data,
        session: localStorage.getItem('Session')
    }, {
        headers: new AxiosHeaders("Access-Control-Allow-Origin: *"),
    })
}

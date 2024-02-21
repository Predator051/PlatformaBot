import axios, {AxiosHeaders} from "axios";

export async function SendPost(url, data) {
    return await axios.post("http://localhost:8080/" + url, {
        ...data,
        session: localStorage.getItem('Session')
    }, {
        headers: new AxiosHeaders("Access-Control-Allow-Origin: *"),
    })
}

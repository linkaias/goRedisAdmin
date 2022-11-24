//设置token
export const SetToken = (token) => {
    localStorage.setItem("token", token);
}

export const GetToken = () => {
    return localStorage.getItem("token");
}

export const RemoveToken = () => {
    localStorage.removeItem("token");
}

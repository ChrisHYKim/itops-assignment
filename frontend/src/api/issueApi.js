import axios from axios

const API_BACK_URL ='http://localhost:8080'

const api =axios.create({
    baseURL: API_BACK_URL,
    headers: {
        'Content-Type': 'appliocation/json',
    }
});
// API 호출 준비
const handlerApiErr =(error) =>{
    let errorMsg = '';
    if(error.reponse){
        // error msg 출력 
        errorMsg = error.response.data.error || `Error: ${error.response.status}`;
        throw new Error(errorMsg);
    }
    else if(error.request) {
        errorMsg = "No Respone Error";
    }
    else {
        errorMsg =  `Error setting up request: ${error.message}`;
        
    }
    throw new Error(errorMsg);
};

// API Function FetchAPI (GET,POST)
export const fetchIssues = async (statusFilter = '') => {
    try {
        const rep = await api.get("/issues",{
            params: { status: statusFilter}
        });
        return response.data.issues;
    } catch (error) {
        handlerApiErr(error)
    }
};
// Issue id 조회
export const fetchIssueById = async (id) => {
    try {
        const resp = await api.get('/issue/${id}');
        return resp.data;
    } catch (error) {
        handlerApiErr(error);
    }
};
export const createIssue = async (issueData) =>{
    try {
        const resp = await api.post("/issue", issueData);
        return resp.data;
    } catch (error) {
        handlerApiErr(error);
    }
}
export const updateIssue = async (id, updateData) => {
    try {
        const resp = await api.patch(`/issue/${id}`, updateData);
        return resp.data;
    } catch (error) {
        handlerApiErr(error);
    }
};
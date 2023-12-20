class NinjaUserBasicInfo{
    constructor(addr, ethAddr, twId, createAt) {
        this.address =addr;
        this.eth_addr = ethAddr;
        this.tw_id = twId;
        this.create_at = createAt;
    }
}

function sesDbKeyForSignData(ethAddr){
    return "__session_database_key_for_sign_data__:"+ethAddr
}
function sesDbKeyForNjUserData() {
    return "__session_database_key_for_ninja_user_current__"
}
function setDataFromSessionDB(key, sign_data){
    sessionStorage.setItem(key, JSON.stringify(sign_data));
}

function getDataFromSessionDB(key){
    const storedValue = sessionStorage.getItem(key);
    return storedValue ? JSON.parse(storedValue) : null;
}
function clearSessionStorage() {
    sessionStorage.clear();
}

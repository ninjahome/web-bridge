class NinjaUserBasicInfo {
    constructor(addr, ethAddr, twId, createAt) {
        this.address = addr;
        this.eth_addr = ethAddr;
        this.tw_id = twId;
        this.create_at = createAt;
    }

    static syncToSessionDbForApiResponse(response) {
        const ninjaObj = JSON.parse(response)
        if (!ninjaObj.eth_addr){
            throw new Error("invalid ninja user info")
        }
        setDataToSessionDB(sesDbKeyForCurrentUserEthAddr(), ninjaObj.eth_addr);
        sessionStorage.setItem(sesDbKeyForNjUserData(ninjaObj.eth_addr), response);
        return ninjaObj;
    }

    static loadCurrentNJUserObj() {
        const curUsrEthAddr = getDataFromSessionDB(sesDbKeyForCurrentUserEthAddr())
        const savedUserInfo = getDataFromSessionDB(sesDbKeyForNjUserData(curUsrEthAddr))
        if (!savedUserInfo) {
            return null;
        }
        return new NinjaUserBasicInfo(savedUserInfo.address, savedUserInfo.eth_addr,
            savedUserInfo.tw_id, savedUserInfo.create_at);
    }
}

class TwitterBasicInfo {
    constructor(id, name, username, avatarUrl, bio) {
        this.id = id;
        this.name = name;
        this.username = username;
        this.profile_image_url = avatarUrl;
        this.description = bio;
    }
    static loadTwBasicInfo(TwitterID){
       const storedData=  getDataFromSessionDB(sesDbKeyForTwitterUserData(TwitterID))
        if (!storedData){
            return null
        }
        return new TwitterBasicInfo(storedData.id,storedData.name, storedData.username,
            storedData.profile_image_url, storedData.description);
    }
    static cacheTwBasicInfo(objStr){
        const obj = JSON.parse(objStr)
        if (!obj.id){
            throw new Error("invalid twitter basic info")
        }
        sessionStorage.setItem(sesDbKeyForTwitterUserData(obj.id), objStr);
        return obj;
    }
}
function sesDbKeyForTwitterUserData(TwitterID) {
    return "__session_database_key_for_twitter_user_data__:" + TwitterID
}
function sesDbKeyForNjUserData(ethAddr) {
    return "__session_database_key_for_ninja_user_data__:" + ethAddr
}

function sesDbKeyForCurrentUserEthAddr() {
    return "__session_database_key_for_ninja_user_current_address__"
}

function setDataToSessionDB(key, sign_data) {
    sessionStorage.setItem(key, JSON.stringify(sign_data));
}

function getDataFromSessionDB(key) {
    const storedValue = sessionStorage.getItem(key);
    return storedValue ? JSON.parse(storedValue) : null;
}

function clearSessionStorage() {
    sessionStorage.clear();
}

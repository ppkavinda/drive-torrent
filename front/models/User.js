export default class User {
    constructor(id, name, imageURL, email) {
        this.ID = id
        this.DisplayName = name
        this.ImageURL = imageURL
        this.Email = email
    }
    isLoggedIn () {
        return this.ID != undefined
    }
}
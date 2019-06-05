class Media {
    constructor(title) {
        this._title = title;
        this._isCheckedOut = false;
        this._ratings = [];
    }
    get title() {
        return this._title;
    }
    get isCheckedOut() {
        return this._isCheckedOut;
    }
    get ratings() {
        return this._ratings;
    }
    set isCheckedOut(x) {
        this._isCheckedOut = x;
    }
    toggleCheckOutStatus() {
        if (this._isCheckedOut === true) {
            this._isCheckedOut = false;
        } else {
            this._isCheckedOut = true;
        }
    }
    getAverageRating() {
        let reducer = function add(sum, item) {
            return sum + item;
        };
        let total = this._ratings.reduce(reducer,0);
        return total;
    }
    addRating(xx) {
        this._ratings.push(xx);
    }
}

class Book extends Media {
    constructor(title, author, pages) {
        super(title);
        this._author = author;
        this._pages = pages;
    }
    get author() {
        return this._author;
    }
    get pages() {
        return this._pages;
    }
}

class Movie extends Media {
    constructor(title, director, runTime) {
        super(title);
        this._director = director;
        this._runTime = runTime;
    }
    get director() {
        return this._director;
    }
    get runTime() {
        return this._runTime;
    }
}

class CD extends Media {
    constructor(title, artist, songs) {
        super(title);
        this._artist = artist;
        this._songs = songs;
    }
    get artist() {
        return this._artist;
    }
    shuffle() {
        const randomSort = () => {
            return Math.random() > 0.5 ? -1 : 1;
        }
        return this._songs.sort(randomSort);
    }
}

const historyOfEverything = new Book("A Short History of Nearly Everything", "Bill Bryson", 544);
historyOfEverything.toggleCheckOutStatus();
console.log(historyOfEverything.isCheckedOut);
historyOfEverything.addRating(4);
historyOfEverything.addRating(5);
historyOfEverything.addRating(5);
console.log(historyOfEverything.getAverageRating());

const speed = new Movie("Speed", "Jan de Bont", 116);
speed.toggleCheckOutStatus();
console.log(speed.isCheckedOut);
speed.addRating(1);
speed.addRating(1);
speed.addRating(5);
console.log(speed.getAverageRating());

const newCD = new CD("Fashion", "Micheal", ["I love you more than I can say", "Stars", "go to Mars"]);
console.log(newCD.shuffle());
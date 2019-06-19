<template>
    <div class="container">
        <div class="row">
            <div class="col s4">
                <div class="card-image waves-effect waves-block waves-light">
                    <img class="activator" :src=movie.medium_cover_image>
                </div>
            </div>
            <div class="col s8">
                <div class="card">
                    <div class="card-content">
                    <div class="card-title"><h4>{{movie.title}}</h4></div>
                        <p>{{movie.year}}</p><br>
                        <div class="row">
                            <div class="col s2">
                                <img src="https://yts.am/assets/images/website/logo-imdb.svg">
                            </div>
                            <div class="col s2">
                                <p>{{movie.rating}}</p>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col" v-for="(tag, index) in movie.genres" :key="index">
                                <span class="btn orange lighten-4 black-text">{{tag}}</span>
                            </div>
                        </div>
                        <br>
                            <p>synopsis: {{movie.synopsis}} </p>
                        <br>
                    </div>
                </div>
            </div>
        </div>
        <div class="card" v-for="(torrent, index) in movie.torrents" :key="index">
            <div class="card-content">
                <div class="card-title"><h5>{{torrent.type}} {{torrent.quality}} <small>({{torrent.size}})</small></h5>
                    <a class="waves-effect waves-light btn-small" :href="torrent.url">Torrent</a>
                    <a class="waves-effect waves-light btn-small" :href="getMagnet(movie.title_long, torrent.hash)">Magnent</a>
                    <a class="btn orange waves-effect waves-light " @click="downloadTorrent(movie.title_long, torrent)">Drive torrent</a>
                </div>
                <ul ref="coll" class="collapsible">
                    <li>
                    <div class="collapsible-header"><i class="material-icons">filter_drama</i>Drive Links</div>
                    <div class="collapsible-body">
                        <table class="striped">
                            <tr>
                                <th>File</th>
                                <th>Up votes</th>
                                <th>Down votes</th>
                                <th title="does the link work or not?">Works ?</th>
                            </tr>
                            <tr  v-for="(file ,name, i) in links[torrent.hash.toLowerCase()]" :key="i">
                                <td><a :href="file.link">{{ name }}</a></td>
                                <td>{{ file.upvotes}}</td>
                                <td>{{ file.downvotes }}</td>
                                <td>
                                    <button class="btn" @click="upvote(torrent.hash, name)"><i class="material-icons">check</i></button>
                                    <button class="btn red" @click="downvote(torrent.hash, name)"><i class="material-icons">close</i></button>
                                </td>
                            </tr>
                            <tr v-if="!links[torrent.hash]">No Drive links yet</tr>
                        </table> 
                    </div>
                    </li>
                </ul>
            </div>
        </div>
        <div class="row">
            <iframe type="text/html" 
                :src="'https://youtube.com/embed/'+movie.yt_trailer_code+'?autoplay=0'" 
                frameborder="0" width="100%" height="360">
            </iframe>
        </div>
    </div>
</template>

<script>
import router from '../router'
import { db } from '../config/firebase'
import { firebase } from "@firebase/app"

export default {
    name: 'movie',
    data(){
        return{
            id: 0,
            movie:{},
            links: {}
        }
    },
    methods: {
        downloadTorrent(movieName, torrent) {
            let magnet = this.getMagnet(movieName, torrent.hash);

            axios.post("/new/magnet", { magnet })
            .then(res => {
                this.$router.push({name: 'downloading'})
            })
                .catch(err => {
                if (err.response.status == 401) window.location.replace('/login')
            });
        },
        upvote(hash, name) {
            db.doc(`hashes/${hash}`).set({[name]: {upvotes: firebase.firestore.FieldValue.increment(1)}}, {merge: true})
        },
        downvote(hash, name) {
            db.doc(`hashes/${hash}`).set({[name]: {downvotes: firebase.firestore.FieldValue.increment(1)}}, {merge: true})
        },
        getFilesOfHash(hash) {
            db.doc(`hashes/${hash}`).onSnapshot(snap  => {
                const data = snap.data()
                if (!data) return
                this.$set(this.links, hash, data)
            })
            
        },
        getMagnet(movieName, infoHash) {
            return `magnet:?xt=urn:btih:${infoHash}&dn=${encodeURI(movieName)}&tr=udp://glotorrents.pw:6969/announce&tr=udp://tracker.opentrackr.org:1337/announce&tr=udp://torrent.gresille.org:80/announce&tr=udp://tracker.openbittorrent.com:80&tr=udp://tracker.coppersurfer.tk:6969&tr=udp://tracker.leechers-paradise.org:6969&tr=udp://p4p.arenabg.ch:1337&tr=udp://tracker.internetwarriors.net:1337`
        }
    },
    mounted(){
        setTimeout(() => {
            var elems = document.querySelectorAll('.collapsible');
            var instances = M.Collapsible.init(elems);
        }, 2000);

        this.id = this.$route.params.id;
        db.collection('films').doc(`${this.id}`).get().then((querySnapshot) => {
            this.movie = querySnapshot.data();
            this.movie.torrents.forEach(torrent => {
                this.getFilesOfHash(torrent.hash.toLowerCase())
            });
        })
    }
}
</script>

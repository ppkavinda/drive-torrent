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
        <div class="card">
            <div class="card-content">
                <div class="row valign-wrapper" v-for="(torrent, index) in movie.torrents" :key="index">
                    <div class="col l4">
                        Quality:{{torrent.type}} {{torrent.quality}}
                    </div>
                    <div class="col l2">
                        Size: {{torrent.size}}
                    </div>
                    <div class="col l2">
                        <a class="waves-effect waves-light btn-small" :href="torrent.url">Torrent</a>
                    </div>
                    <div class="col l2">
                        <a class="waves-effect waves-light btn-small" :href="getMagnet(movie.title_long, torrent.hash)">Magnent</a>
                    </div>
                    <div class="col l2">
                        <!-- <a class="waves-effect waves-light btn-small">Direct</a> -->
                    <button class="btn orange waves-effect waves-light btn-large right" @click="downloadTorrent(movie.title_long, torrent)">Drive torrent</button>
                    </div>
                </div>
            </div>
        </div>
        <div class="card">
            <div class="card-content">
                <table class="striped">
                    <tr>
                        <th>Link</th>
                        <th>Quality</th>
                        <th>Up votes</th>
                        <th>Down votes</th>
                        <th>Alive</th>
                    </tr>
                    <tr>
                        <td><a href="#">Sample Google Drive1</a></td>
                        <td>720p</td>
                        <td>105</td>
                        <td>20</td>
                        <td>Alive</td>
                    </tr>
                    <tr>
                        <td><a href="#">Sample Google Drive2</a></td>
                        <td>1080p</td>
                        <td>101</td>
                        <td>200</td>
                        <td>Dead</td>
                    </tr>
                </table> 
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
import { db } from '../config/firebase';

export default {
    name: 'movie',
    data(){
        return{
            id: 0,
            movie:{}
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
        getMagnet(movieName, infoHash) {
            return `magnet:?xt=urn:btih:${infoHash}&dn=${encodeURI(movieName)}&tr=udp://glotorrents.pw:6969/announce&tr=udp://tracker.opentrackr.org:1337/announce&tr=udp://torrent.gresille.org:80/announce&tr=udp://tracker.openbittorrent.com:80&tr=udp://tracker.coppersurfer.tk:6969&tr=udp://tracker.leechers-paradise.org:6969&tr=udp://p4p.arenabg.ch:1337&tr=udp://tracker.internetwarriors.net:1337`
        }
    },
    mounted(){
        this.id = this.$route.params.id;
        db.collection('films').doc(`${this.id}`).get().then((querySnapshot) => {
            this.movie = querySnapshot.data();
        })
    }
}
</script>

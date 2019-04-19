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
                    <div class="card-title">{{movie.title}}</div>
                        <p>{{movie.year}}</p><br>
                        <div class="row">
                            <div class="col" v-for="(tag, index) in movie.genres" :key="index">
                                <span class="btn">{{tag}}</span>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col s2">
                                <img src="https://yts.am/assets/images/website/logo-imdb.svg">
                            </div>
                            <div class="col s2">
                                <p>{{movie.rating}}</p>
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
                    <div class="col s4">
                        Quality:{{torrent.type}} {{torrent.quality}}
                    </div>
                    <div class="col s2">
                        Size: {{torrent.size}}
                    </div>
                    <div class="col s2">
                        <a class="waves-effect waves-light btn-small" :href="torrent.url">Torrent</a>
                    </div>
                    <div class="col s2">
                        <a class="waves-effect waves-light btn-small" :href="'magnet:?xt=urn:btih:'+torrent.hash+'&dn=Url+Encoded+Movie+Name&tr=udp://glotorrents.pw:6969/announce&tr=udp://tracker.opentrackr.org:1337/announce&tr=udp://torrent.gresille.org:80/announce&tr=udp://tracker.openbittorrent.com:80&tr=udp://tracker.coppersurfer.tk:6969&tr=udp://tracker.leechers-paradise.org:6969&tr=udp://p4p.arenabg.ch:1337&tr=udp://tracker.internetwarriors.net:1337'">Magnent</a>
                    </div>
                    <div class="col s2">
                        <a class="waves-effect waves-light btn-small">Direct</a>
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
                        <td><a href="https://google.com">Sample Google Drive1</a></td>
                        <td>720p</td>
                        <td>105</td>
                        <td>20</td>
                        <td>Alive</td>
                    </tr>
                    <tr>
                        <td><a href="https://google.com">Sample Google Drive2</a></td>
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
    mounted(){
        this.id = this.$route.params.id;
        db.collection('films').doc(`${this.id}`).get().then((querySnapshot) => {
            this.movie = querySnapshot.data();
        })
    }
}
</script>

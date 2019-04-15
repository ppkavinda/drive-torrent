<template>
    <div class="row">
        <div v-for="(film, index) in films" :key="index" class="col s3">
          <div class="card sticky-action">
            <div class="card-image waves-effect waves-block waves-light">
              <img class="activator" :src="film.medium_cover_image">
            </div>
            <div class="card-content">
              <span class="card-title activator grey-text text-darken-4">{{film.title}}<i class="material-icons right">more_vert</i></span>
              <p>{{film.year}}</p>
              <p><router-link :to="'/movie/'+film.id"> Download</router-link></p>
            </div>
            <div class="card-reveal">
              <span class="card-title grey-text text-darken-4">{{film.title}}<i class="material-icons right">close</i></span>
              <p>
                <b>IMDB:</b> <a v-bind:href="'https://www.imdb.com/title/'+film.imdb_code">{{film.rating}}</a><br>
                <b>Run time:</b> {{film.runtime}}<br>
                <b>Summary:</b><br>
                {{film.synopsis}}
              </p>
            </div>
          </div>
        </div>
    </div>
</template>

<script>
import {db} from '../config/firebase';

export default {
    data(){
      return {
        films: [],
        ref: db.collection('films')
      }
    },
    mounted(){
      db.collection('films').orderBy('date_uploaded','desc').limit(9).get().then((querySnapshot)=> {
        querySnapshot.forEach((doc) => {
          this.films.push(doc.data());
        })
      })
    }

}
</script>


<template>
    <div>
        <h3>Recent YTS</h3>
        <ais-instant-search :search-client="searchClient" index-name="films">
            <!-- <ais-search-box placeholder="Search YTS torrents" /> -->
            <ais-search-box >
                <div slot-scope="{ currentRefinement, isSearchStalled, refine }">
                    <div class="row">
                      <div class="input-field ">
                        <i class="material-icons prefix">search</i>
                        <input
                            type="text"
                            v-model="currentRefinement"
                            @input="refine($event.currentTarget.value)"
                            id="searchYts"
                        />
                        <label for="searchYts">Search YTS torrents</label>
                        <span :hidden="!isSearchStalled">Loading...</span>
                      </div>
                    </div>
                </div>
            </ais-search-box>
            <ais-hits>
                <div class="row" slot-scope="{ items }">
                    <div v-for="(film, index) in items" :key="index" class="col s12 m4 l3">
                      <div class="card sticky-action hoverable">
                        <div class="card-image waves-effect waves-block waves-light">
                          <img class="activator" :src="film.medium_cover_image">
                        </div>
                        <div class="card-content">
                          <span class="card-title activator grey-text text-darken-4 truncate">
                            <i class="material-icons right">more_vert</i>
                            {{film.title}}
                          </span>
                          <p><strong>Year:</strong> {{film.year}}</p>
                          <strong>IMDB:</strong> <a :href="'https://www.imdb.com/title/'+film.imdb_code">{{film.rating}}</a><br>
                        </div>
                        <div class="card-action">
                          <router-link :to="'/movie/'+film.id"> Download</router-link>
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
            </ais-hits>
            <ais-pagination />
        </ais-instant-search>
    </div>
</template>

<script>
import algoliasearch from "algoliasearch/lite";
// import "instantsearch.css/themes/algolia-min.css";

export default {
    data() {
        return {
            searchClient: algoliasearch(
                "GY5MI8OEXN",
                "fc927ba5838553079e05db0a39d47fcd"
            )
        };
    }
};
</script>

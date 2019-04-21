<template>
    <ais-instant-search :search-client="searchClient" index-name="films">
        <ais-search-box />
        <ais-hits>
            <div slot="item" slot-scope="{ item }">
                <div class="card-image waves-effect waves-block waves-light">
                    <router-link :to="'/movie/'+item.id">
                        <img class="activator" :src="item.medium_cover_image">
                    </router-link>
                </div>
                <div class="card-content">
                    <router-link :to="'/movie/'+item.id"><span class="card-title activator grey-text text-darken-4">{{item.title}}<i class="material-icons right">more_vert</i></span></router-link>
                    <p>{{item.year}}</p>
                    <p><router-link :to="'/movie/'+item.id"> Download</router-link></p>
                </div>
            </div>
        </ais-hits>
        <ais-pagination />
    </ais-instant-search>
</template>

<script>
import algoliasearch from 'algoliasearch/lite';
import 'instantsearch.css/themes/algolia-min.css';

export default {
    data() {
        return {
            searchClient: algoliasearch(
            'GY5MI8OEXN',
            'fc927ba5838553079e05db0a39d47fcd'
            ),
        };
  }
}
</script>

<style>
body {
  font-family: sans-serif;
  padding: 1em;
}
</style>

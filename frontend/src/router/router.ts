import { createWebHashHistory, createRouter } from 'vue-router'

import HomeView from '../views/StartScreen.vue'
import MovieListView from '../views/MovieListView.vue'
import ShowListView from '../views/ShowListView.vue'
import MovieDetailsView from '../views/MovieDetailsView.vue'
import MovieStreamView from '../views/MovieStreamView.vue'
import ShowDetailsView from '../views/ShowDetailsView.vue'
import ShowStreamView from '../views/ShowStreamView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/movies', component: MovieListView },
  { path: '/movies/:id', component: MovieDetailsView },
  { path: '/movies/stream/:id', component: MovieStreamView },
  { path: '/series', component: ShowListView },
  { path: '/series/:id', component: ShowDetailsView },
  { path: '/series/stream/:id/:season/:episode/:episodeAmt', component: ShowStreamView },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
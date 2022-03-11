<template>
  <v-app teal>
    <NavBar v-show="currentUser" />
    <v-content>
      <router-view />
    </v-content>
    <v-footer color="blue darken-4" app>
      <span class="white--text">IT Elektrikal © 2021</span>
    </v-footer>
  </v-app>
</template>
 
<script>
import NavBar from "./components/NavBar";
import { authenticationService } from '@/_services';
export default {
  name: "App",
  components: {
    NavBar,
  },
  data () {
    return {
      currentUser: null
    };
  },
  created () {
    authenticationService.currentUser.subscribe(x => this.currentUser = x);
  },
};
</script>
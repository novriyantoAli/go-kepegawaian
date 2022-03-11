<template>
  <section>
    <v-container class="pa-2" fluid>
      <v-row>
        <v-col>
          <v-alert v-if="errors" type="error" class="mt-4">{{errors}}</v-alert>
          <v-card color="#385F73" dark v-if="pegawaiResult">
            <v-card-text class="white--text">
              <div class="headline mb-2">{{currentUser.nik}}</div>
              Sepertinya profil anda belum lengkap, segera lengkapi data profile untuk keperluan kepegawaian
            </v-card-text>
 
            <v-card-actions>
              <v-btn text>Lengkapi Sekarang!</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </section>
</template>

<script>
import { authenticationService } from '@/_services';
export default {
	data () {
		return {
			currentUser: null,
			errors: null,
			loading: false,
			pegawaiResult: false
		};
	},
	created () {
		authenticationService.currentUser.subscribe(x => this.currentUser = x);
	},
	mounted: function () {
		this.loading = true;
		
	},
}
</script>
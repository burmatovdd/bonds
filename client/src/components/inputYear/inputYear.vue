<template>
	<div class="container">
		<div class="title">Enter Year</div>
		<input id="input" type="number" class="input" ref="yearInput" />
		<button id="button" class="button" @click="send">Submit</button>
		<div  :class="{overlay: !active}">
			<table class="bondsTable">
				<caption align="bottom">
					Total:
					{{
						total
					}}
				</caption>
				<thead>
					<tr>
						<th class="thName">Name</th>
						<th class="thCount">Count</th>
						<th class="thDate" :colspan="months.length">Date</th>
					</tr>
					<tr v-if="months.length !== 0">
						<th colspan="2" class="empty"></th>
						<th>
							<table class="months">
								<tr id="months" class="months-day">
									<th v-for="month in months" :key="month">
										{{ month }}
									</th>
								</tr>
							</table>
						</th>
					</tr>
				</thead>
				<tbody>
					<tr :key="item.name" id="bonds" v-for="item in response.allInfos.bondInfos" class="bondsTr">
						<td>
							{{ item.bond.name }}
						</td>
						<td>
							{{ item.bond.count }}
						</td>
						<td>
							<table style="width: 100%; table-layout: fixed">
								<tr>
									<template v-for="month in months">
										<td
											:key="`${month}-if`"
											v-if="item.coupons.find((coupon) => month === coupon.date)"
										>
											{{ item.coupons.find((coupon) => month === coupon.date).value }}
										</td>

										<td :key="`${month}-else`" v-else></td>
									</template>
								</tr>
							</table>
						</td>
						<td>
							<button class="button" @click="_delete(item)">Delete</button>
						</td>
					</tr>
				</tbody>
				<tfoot>
					<tr>
						<td></td>
						<td>Total:</td>
						<td>
							<table style="width: 100%; table-layout: fixed">
								<tr>
									<template v-for="month in months">
										<td :key="`${month}-if`" v-if="monthTotal(month) !== 0">
											{{ monthTotal(month) }}
										</td>
										<td :key="`${month}-else`" v-else></td>
									</template>
								</tr>
							</table>
						</td>
					</tr>
				</tfoot>
			</table>
		</div>
	</div>
</template>

<script>
import { defineComponent, ref } from "vue";
import *as storage from "../../storage";
import *as httpClient from "../../httpClient";
import {useRouter} from "vue-router";



export default defineComponent({
	setup() {
		const yearInput = ref();
		return {
			yearInput,
		};
	},
	data() {
		return {
			months: [],
			response: {
                allInfos: {
                    bondInfos: [],
                    months: [],
                },
			},
            active: false
		};
	},
	computed: {
		total() {
			if (this.$data.response) {
				const total = this.$data.response.allInfos.bondInfos.reduce((accumulator, bond) => {
					return (
						accumulator +
						bond.coupons.reduce((accumulator, coupon) => {
							return accumulator + coupon.value;
						}, 0)
					);
				}, 0);

				return total.toFixed(2);
			}

			return 0;
		},
	},
	methods: {
		monthTotal(month) {
			const total = this.$data.response.allInfos.bondInfos.reduce((accumulator, bond) => {
				return (
					accumulator +
					bond.coupons.reduce((accumulator, coupon) => {
						return coupon.date === month ? accumulator + coupon.value : accumulator;
					}, 0)
				);
			}, 0);
            return total.toFixed(2);
		},
		async _delete(bond) {
            this.$data.response.allInfos.bondInfos = this.$data.response.allInfos.bondInfos.filter((bondObj) => bondObj.bond.name !== bond.bond.name);
            // console.log("bondName: ", bond.name)

            let sendUrl = "/api/delete";

            let postInfo = httpClient.Post(sendUrl,{name: bond.bond.name});
		},

		async send() {
            if (this.active !== true){
                this.active = !this.active;
            }
            this.$data.months.length = 0; // Clear months array every time we click send button
            const year = this.$refs.yearInput.value;

            for (let i = 1; i <= 12; i++) {
                this.$data.months.push(`${year}-${i <= 9 ? `0${i}` : i}`);
            }
            let token = storage.get("token");
            if (token == null){
                await this.$router.push('/');
            }

            let sendUrl = "/api/year";

            let postInfo = httpClient.Post(sendUrl,{year: this.$refs.yearInput.value});
            console.log("postInfo: ", postInfo);
            postInfo.then((data) => {
                this.$data.response = data;
            })

		},
	},
});
</script>

<style lang="css">
@import "../../style/style.css";
@import "./inputYear.css";
</style>

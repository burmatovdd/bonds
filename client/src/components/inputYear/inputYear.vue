<template>
	<div class="container">
		<div class="title">Enter Year</div>
		<input id="input" type="number" class="input" ref="yearInput" />
		<button id="button" class="button" @click="send">Submit</button>
		<div class="table-responsive">
			<table class="bondsTable">
				<caption align="bottom">
					ИТОГО:
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
					<tr :key="bond.name" id="bonds" v-for="bond in response.bonds" class="bondsTr">
						<td>
							{{ bond.name }}
						</td>
						<td>
							{{ bond.count }}
						</td>
						<td>
							<table style="width: 100%; table-layout: fixed">
								<tr>
									<template v-for="month in months">
										<td
											:key="`${month}-if`"
											v-if="bond.coupons.find((coupon) => month === coupon.date)"
										>
											{{ bond.coupons.find((coupon) => month === coupon.date).value }}
										</td>

										<td :key="`${month}-else`" v-else></td>
									</template>
								</tr>
							</table>
						</td>
						<td>
							<button class="button" @click="_delete(bond)">Delete</button>
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
				bonds: [],
				months: [],
			},
		};
	},
	computed: {
		total() {
			if (this.$data.response) {
				const total = this.$data.response.bonds.reduce((accumulator, bond) => {
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
			const total = this.$data.response.bonds.reduce((accumulator, bond) => {
				return (
					accumulator +
					bond.coupons.reduce((accumulator, coupon) => {
						return coupon.date === month ? accumulator + coupon.value : accumulator;
					}, 0)
				);
			}, 0);
			return total;
		},
		async _delete(bond, url = "http://localhost:8080/delete") {
			this.$data.response.bonds = this.$data.response.bonds.filter((bondObj) => bondObj.name !== bond.name);

			// await fetch(url, {
			// 	method: "POST",
			// 	headers: {
			// 		"Content-Type": "application/json",
			// 	},
			// 	body: JSON.stringify(bond.name),
			// });
		},

		async send(url = "http://localhost:8080/year") {
			// await fetch(url, {
			// 	method: "POST",
			// 	headers: {
			// 		"Content-Type": "application/json",
			// 	},
			// 	body: JSON.stringify(tmpl),
			// })
			// 	.then((response) => {
			// 		return response.json();
			// 	})
			// 	.then((data) => {
			// 		this.$data.response = data;
			// 	})
			// 	.catch(console.error);
			this.$data.response = {
				bonds: [
					{
						name: "RU000A104FG2",
						count: 2,
						coupons: [
							{
								date: "2022-04",
								value: 54.36,
							},
							{
								date: "2022-07",
								value: 54.36,
							},
							{
								date: "2022-10",
								value: 54.36,
							},
						],
					},
					{
						name: "RU000A1040V2",
						count: 2,
						coupons: [
							{
								date: "2022-02",
								value: 62.32,
							},
							{
								date: "2022-05",
								value: 62.32,
							},
							{
								date: "2022-08",
								value: 62.32,
							},
							{
								date: "2022-11",
								value: 62.32,
							},
						],
					},
					{
						name: "RU000A104CJ3",
						count: 5,
						coupons: [
							{
								date: "2022-03",
								value: 152.7,
							},
							{
								date: "2022-06",
								value: 152.7,
							},
							{
								date: "2022-09",
								value: 152.7,
							},
							{
								date: "2022-12",
								value: 152.7,
							},
						],
					},
				],
				months: [
					{
						date: "02",
						value: 62.32,
					},
					{
						date: "03",
						value: 152.7,
					},
					{
						date: "04",
						value: 54.36,
					},
					{
						date: "05",
						value: 62.32,
					},
					{
						date: "06",
						value: 152.7,
					},
					{
						date: "07",
						value: 54.36,
					},
					{
						date: "08",
						value: 62.32,
					},
					{
						date: "09",
						value: 152.7,
					},
					{
						date: "10",
						value: 54.36,
					},
					{
						date: "11",
						value: 62.32,
					},
					{
						date: "12",
						value: 152.7,
					},
				],
			};

			this.$data.months.length = 0; // Clear months array every time we click send button
			const year = this.$refs.yearInput.value;

			for (let i = 1; i <= 12; i++) {
				this.$data.months.push(`${year}-${i <= 9 ? `0${i}` : i}`);
			}
		},
	},
});
</script>

<style lang="css">
@import "../../style/style.css";
@import "./inputYear.css";
</style>

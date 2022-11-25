<script lang="ts">
	import crown from '../../assets/images/icons-medieval-crown.png'
	import AdminPanel from './AdminPanel.svelte'
	import TeamAction from './TeamAction.svelte'
	import TeamCard from './TeamCard.svelte'
	import Swal from 'sweetalert2'
	import 'sweetalert2/dist/sweetalert2.min.css'
	import { axios } from '../../utils/api'
	import { onMount } from 'svelte'
	let teams: any[] = []
	let update: any = []

	function fetchInfo() {
		axios.get('/am/info').then((res) => {
			teams = res.data?.data
			console.log(teams)
			update = teams.map((e) => null)
		})
	}

	onMount(() => {
		fetchInfo()
	})

	function onScoreChanged(id: number, score: number) {
		console.log(`score changed: ${id} ${score}`)
		update[id] = score
		console.log(update)
	}

	function onSubmit() {
		if (update.length !== teams.length) {
			Swal.fire({
				timer: 2000,
				icon: 'error',
				title: 'Error',
				text: 'No changes made',
			})
			return
		}
		axios
			.patch('/am/score', { update })
			.then((res) => {
				if (res.data?.success) {
					Swal.fire({
						timer: 2000,
						icon: 'success',
						title: 'Success',
						text: res.data?.message,
					}).then(() => {
						update.length = 0
						console.log('clear update')
					})
				} else {
					Swal.fire({
						timer: 2000,
						icon: 'error',
						title: 'Error',
						text: res.data?.message,
					})
				}
				fetchInfo()
			})
			.catch((res) => {
				Swal.fire({
					timer: 2000,
					icon: 'error',
					title: 'Error',
					text: 'Something went wrong!',
				})
			})
	}

	$: submitDisabled =
		update.length !== teams.length || update.some((e) => e === null)

	function pause() {
		return () => {
			axios
				.patch('/am/card/pause')
				.then((res) => {
					if (res.data?.success) {
						Swal.fire({
							timer: 2000,
							icon: 'success',
							title: 'Success',
							text: res.data?.message,
						}).then(() => {
							update.length = 0
							console.log('clear update')
						})
					} else {
						Swal.fire({
							timer: 2000,
							icon: 'error',
							title: 'Error',
							text: res.data?.message,
						})
					}
					fetchInfo()
				})
				.catch((res) => {
					Swal.fire({
						timer: 2000,
						icon: 'error',
						title: 'Error',
						text: 'Something went wrong!',
					})
				})
		}
	}
</script>

<div
	class="w-full h-full flex items-stretch "
	style="background-color: rgb(248, 251, 253)"
>
	<div class="fixed left-0 bottom-0 p-2">
		<button
			on:click={pause()}
			class="p-2 rounded-lg bg-gray-900 hover:bg-gray-800 font-semi text-white"
			>Play/Pause
		</button>
	</div>
	<div
		class="w-[20vw] h-full bg-white flex flex-col"
		style="box-shadow: rgb(232 232 232 / 50%) 0px -2px 2px; border-right: 1px solid rgba(217, 217, 217,.4)"
	>
		<div
			class="h-[64px] flex items-center justify-center font-bold text-lg gap-2"
			style="border-bottom: 1px solid rgb(217, 217, 217); "
		>
			<img width={32} height={32} src={crown} alt="icons-crown" /> League Manager
		</div>

		<div class="flex-1 flex flex-col">
			{#each teams as team}
				<TeamCard
					id={team?.id}
					label={team?.name}
					name={team?.school}
				/>
			{/each}
		</div>
	</div>
	<div class="flex-1 h-full flex flex-col">
		<div
			class="flex-1 w-[calc(80vw-6rem)]"
			style="overflow-y: hidden; overflow-x: scroll;"
		>
			<div class="flex flex-row flex-nowrap">
				{#if teams.length > 0}
					{#each teams[0].scores as team, index}
						<AdminPanel round={index + 1} {teams} />
					{/each}
				{/if}
			</div>
		</div>
	</div>
	<div
		class="w-[6rem] h-full bg-white flex flex-col"
		style="box-shadow: rgb(232 232 232 / 50%) 0px -2px 2px;  border-left: 1px solid rgba(217, 217, 217,.4)"
	>
		<div
			class="h-[64px] flex items-center justify-center font-bold text-md gap-2"
			style="border-bottom: 1px solid rgb(217, 217, 217); "
		>
			Round {teams.length > 0 ? teams[0].scores.length + 1 : 1}
		</div>
		<div class="flex-1 flex flex-col">
			{#each teams as team, index}
				<TeamAction
					id={index}
					choose={update[index]}
					{onScoreChanged}
				/>
			{/each}
			<div class="p-2 mt-4">
				<button
					class="p-2 rounded-lg {submitDisabled
						? 'bg-gray-400'
						: 'bg-emerald-500'} text-white"
					on:click={onSubmit}
					disabled={submitDisabled}>Submit</button
				>
			</div>
		</div>
	</div>
</div>

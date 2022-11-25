<script lang="ts">
	import iconsCrystal from '../../assets/images/icons-crystal.png'
	import icons1st from '../../assets/images/icons-1stmedal-64.png'
	import crown from '../../assets/images/icons-medieval-crown.png'
	import up from '../../assets/images/back-icon.png'
	import { axios } from '../../utils/api'
	import Swal from 'sweetalert2'
	import 'sweetalert2/dist/sweetalert2.min.css'
	let show = -1
	let count = 0

	function setShow(n: number) {
		return () => {
			if (show === n) {
				show = -1
				return
			}
			show = n
			onModeChange()
		}
	}

	function increase() {
		count++
	}

	function decrease() {
		count--
	}

	function increment() {
		axios.patch('/am/preview/increment')
	}

	function onModeChange() {
		;[
			() => {
				console.log(`set to mode: `, 'preview')
				axios
					.patch('/am/mode', {
						mode: 'preview',
					})
					.then((res) => {
						Swal.fire({
							timer: 500,
							icon: 'success',
							title: 'Success',
							text: 'Mode changed to preview',
						})
					})
			},
			() => {
				console.log(`set to mode: `, 'started')
				axios
					.patch('/am/mode', {
						mode: 'started',
					})
					.then((res) => {
						Swal.fire({
							timer: 500,
							icon: 'success',
							title: 'Success',
							text: 'Mode changed to started',
						})
					})
			},
			() => {
				console.log(`set to mode: `)
				axios
					.patch('/am/mode', {
						mode: 'ended',
					})
					.then((res) => {
						Swal.fire({
							timer: 500,
							icon: 'success',
							title: 'Success',
							text: 'Mode changed to ended',
						})
					})
			},
		][show]()
	}
</script>

<div class="flex flex-col p-4 gap-6 h-full w-full">
	<div class="flex items-center gap-1 fixed">
		<h1 class="text-7xl font-bold text-white flex-1">Admin Console</h1>
		<div class="flex gap-4">
			<button class="rounded shadow" />
		</div>
	</div>
	<div class="flex-1 flex items-center justify-center gap-4">
		<button
			class="{show === 0
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(0)}
		>
			<img width={32} height={32} src={crown} alt="icons-crown" />
			Show Preview
		</button>
		<button
			class="{show === 1
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(1)}
		>
			<img
				width={32}
				height={32}
				src={iconsCrystal}
				alt="icons-crystal"
			/> Show Leaderboard
		</button>
		<button
			class="{show === 2
				? 'opacity-100'
				: 'opacity-70'} transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-6 text-white bg-slate-800 hover:bg-slate-900 active:bg-slate-900"
			on:click={setShow(2)}
		>
			<img width={32} height={32} src={icons1st} alt="icons-medal" /> Show
			Podium
		</button>

		<div class="w-2 h-12 bg-[rgba(255,255,255,.3)] rounded-full mx-5" />

		<button
			class="transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-5 text-white border-4 border- border-white hover:bg-[rgba(255,255,255,.1)] active:bg-[rgba(255,255,255,.2)]"
			on:click={increment}
		>
			<img
				width={16}
				height={16}
				src={up}
				class="rotate-90"
				alt="icons-medal"
			/> Reveal
		</button>
		<!-- <button
			class="transition-all flex gap-2 font-semibold text-2xl items-center justify-between rounded-2xl shadow cursor-pointer p-5 text-white border-4 border- border-white hover:bg-[rgba(255,255,255,.1)] active:bg-[rgba(255,255,255,.2)]"
			on:click={decrease}
		>
			<img
				width={16}
				height={16}
				src={up}
				class="-rotate-90"
				alt="icons-medal"
			/> Unreveal
		</button> -->
	</div>
</div>

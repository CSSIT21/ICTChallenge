<script lang="ts">
	import { onMount } from 'svelte'

	import iconsTopic1 from '../../assets/images/icons-topic1.png'
	import iconsTopic2 from '../../assets/images/icons-topic2.png'
	import iconsTopic3 from '../../assets/images/icons-topic3.png'
	import iconsTopic4 from '../../assets/images/icons-topic4.png'
	import iconsTopic5 from '../../assets/images/icons-topic5.png'

	import QuestionCard from 'src/lib/components/QuestionCard.svelte'
	import type { Topic, OpenQuestion } from 'src/types/question'

	import Modal from './QuestionModal.svelte'
	import FlippedCard from './FlippedCard.svelte'

	import { onDestroy } from 'svelte'

	let showModal = false
	let cardId: number
	let cardIndex: number = 0
	let icon: string = ''
	let topicColor: string = ''
	let minute = 0
	let sec = 0
	let socketRetrieve: boolean = false
	export let question: Topic
	export let colIndex: number
	export let openQuestion: OpenQuestion
	export let client

	const widthCard: string = '324px'
	const heightCard: string = '180px'
	const widthImg: string = '88px'
	const textSize: string = '40px'
	const iconSize: string = '120px'

	const unsubscribeclient2 = client.subscribe('cd/open', (payload) => {
		openQuestion = payload
		handleOpenModal(payload.question_id, cardIndex)
	})
	const unsubscribeclient3 = client.subscribe('cd/countdown', (payload) => {
		minute = payload.m
		sec = payload.s
		socketRetrieve = true
	})

	onDestroy(() => {
		unsubscribeclient2()
		unsubscribeclient3()
	})

	const handleOpenModal = (id: number, cardIndex: number) => {
		cardId = id
		showModal = true
		cardIndex = cardIndex
	}
	const handleCloseModal = () => {
		showModal = false
	}

	onMount(() => {
		if (colIndex == 0) {
			icon = iconsTopic1
			topicColor = 'topic-color-1'
		} else if (colIndex == 1) {
			icon = iconsTopic2
			topicColor = 'topic-color-2'
		} else if (colIndex == 2) {
			icon = iconsTopic3
			topicColor = 'topic-color-3'
		} else if (colIndex == 3) {
			icon = iconsTopic4
			topicColor = 'topic-color-4'
		} else if (colIndex == 4) {
			icon = iconsTopic5
			topicColor = 'topic-color-5'
		}
	})
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div class="w-[324px] flex flex-col">
	<div class="flex flex-col justify-center h-[150px] mb-4">
		<p
			class="py-2 text-4xl font-semibold text-center {question.title
				.length < 30
				? 'leading-[48px]'
				: 'leading-10'}  {topicColor}"
		>
			{question.title}
		</p>
	</div>
	<div class="h-[620px] flex flex-col justify-between">
		{#each question.cards as card, i (i)}
			<div class={card.opened && 'flipped'}>
				<div class="flip-card">
					<div class="flip-card-inner">
						<div class="flip-card-front">
							<QuestionCard
								img={icon}
								score={card.score}
								textColor={'text-color-' + (colIndex + 1)}
								{widthCard}
								{heightCard}
								{widthImg}
								{textSize}
							/>
						</div>
						<div
							class="flip-card-back"
							style="transform: rotateY(180deg);"
						>
							<FlippedCard {widthCard} {heightCard} {iconSize} />
						</div>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
<Modal
	open={showModal}
	{openQuestion}
	{minute}
	{sec}
	cardCol={colIndex}
	{cardIndex}
	{handleCloseModal}
	{client}
	{socketRetrieve}
/>

<style>
	.flip-card {
		cursor: pointer;
		user-select: none;
	}
	.flip-card-inner {
		width: 100%;
		height: 100%;
		text-align: center;
		transition: transform 0.6s;
		transform-style: preserve-3d;
		box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
	}

	.flipped .flip-card .flip-card-inner {
		transform: rotateY(180deg);
	}

	.flip-card-front,
	.flip-card-back {
		position: absolute;
		width: 100%;
		height: 100%;
		-webkit-backface-visibility: hidden;
		backface-visibility: hidden;
	}
	.topic-color-1 {
		background: -webkit-linear-gradient(0deg, #48a3fa, #5ad2f3);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-2 {
		background: -webkit-linear-gradient(0deg, #ac77e7, #ead7fe);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-3 {
		background: -webkit-linear-gradient(0deg, #ffb34c, #ffe155);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-4 {
		background: -webkit-linear-gradient(0deg, #fc6cb9, #ffc7e5);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-5 {
		background: -webkit-linear-gradient(0deg, #0fbc88, #aff4ee);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>

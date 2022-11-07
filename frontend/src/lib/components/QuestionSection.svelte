<script lang="ts">
	import { onMount } from 'svelte'

	import iconsTopic1 from '../../assets/images/icons-topic1.png'
	import iconsTopic2 from '../../assets/images/icons-topic2.png'
	import iconsTopic3 from '../../assets/images/icons-topic3.png'
	import iconsTopic4 from '../../assets/images/icons-topic4.png'

	import QuestionCard from 'src/lib/components/QuestionCard.svelte'
	import type { Topic } from 'src/types/question'
	import Modal from './QuestionModal.svelte'
	import FlippedCard from './FlippedCard.svelte'

	let showModal = false
	let cardId: number
	let cardIndex: number
	let icon: string = ''
	let topicColor: string = ''
	export let question: Topic
	export let colIndex: number
	export let openQuestion: string
	export let openCard: (cardCol: number, cardIndex: number) => void
	export let getQuestion: (questionId: number) => void

	const widthCard: string = '324px'
	const heightCard: string = '180px'
	const widthImg: string = '88px'
	const textSize: string = '40px'
	const iconSize: string = '120px'

	const handleOpenModal = (id: number, index: number) => {
		cardId = id
		cardIndex = index
		showModal = true
		getQuestion(id)
	}
	const handleCloseModal = (cardCol: number, cardIndex: number) => {
		showModal = false
		openCard(cardCol, cardIndex)
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
		}
	})
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->

<div class="w-[324px] h-[874px] flex flex-col">
	<div class="flex flex-col justify-center h-[150px]">
		<p class="py-2 text-4xl font-semibold text-center {topicColor}">
			{question.title}
		</p>
	</div>
	<div class="h-[600px] flex flex-col justify-between">
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
								cardId={card.id}
								cardIndex={i}
								{handleOpenModal}
							/>
						</div>
						<div
							class="flip-card-back"
							style="transform: rotateY(180deg);"
						>
							<FlippedCard
								{widthCard}
								{heightCard}
								{iconSize}
								cardId={card.id}
								cardIndex={i}
								{handleOpenModal}
							/>
						</div>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
<Modal
	cardCol={colIndex}
	{cardIndex}
	open={showModal}
	{openQuestion}
	{handleCloseModal}
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
</style>

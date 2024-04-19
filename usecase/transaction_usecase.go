package usecase

import (
	"sort"
)

type GoAssestmentUseCaseInterface interface {
	Soal1(nums []int, target int) (result []int, err error)
	Soal2(nums []int) (result [][]int, err error)
	Soal3(s string, words []string) (result []int, err error)
}

type transactionUseCase struct {
}

func NewTransactionUseCase() GoAssestmentUseCaseInterface {
	return &transactionUseCase{}
}

func (t *transactionUseCase) Soal1(nums []int, target int) (result []int, err error) {
	// Membuat map untuk menyimpan nilai dan indeks dari elemen nums
	numMap := make(map[int]int)

	// Melakukan iterasi pada nums
	for i, num := range nums {
		// Mencari selisih yang dibutuhkan untuk mencapai target
		complement := target - num
		// Jika selisih tersebut ada dalam map, artinya kita telah menemukan pasangan yang memenuhi syarat
		if j, ok := numMap[complement]; ok {
			return []int{j, i}, nil
		}
		numMap[num] = i
	}

	return nil, nil
}

func (t *transactionUseCase) Soal2(nums []int) (result [][]int, err error) {
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Menghindari duplikasi untuk nilai i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// Menghindari duplikasi untuk nilai left
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Menghindari duplikasi untuk nilai right
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result, nil
}

func (t *transactionUseCase) Soal3(s string, words []string) (result []int, err error) {
	if len(words) == 0 || len(s) == 0 {
		return result, nil
	}

	wordLength := len(words[0])
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	totalWords := len(wordsMap)

	for i := 0; i <= len(s)-wordLength*totalWords; i++ {
		substrMap := make(map[string]int)
		count := 0

		for j := 0; j < totalWords; j++ {
			word := s[i+j*wordLength : i+(j+1)*wordLength]
			if _, exists := wordsMap[word]; !exists {
				break
			}
			substrMap[word]++
			if substrMap[word] > wordsMap[word] {
				break
			}
			count++
		}

		if count >= totalWords {
			result = append(result, i)
		}
	}

	return result, nil
}

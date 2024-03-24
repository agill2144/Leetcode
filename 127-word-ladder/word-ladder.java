class Solution {
    Set<String> visited = new HashSet<>();
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        //start word can not be end word
        if(beginWord == null || beginWord.length() == 0 || endWord == null || endWord.length() == 0) return 0;

        Queue<String> q = new LinkedList<>();
        int size = 0;
        int tseq = 1;
        q.add(beginWord); 
        visited.add(beginWord);

        while(!q.isEmpty()){
            size = q.size();
            while(size>0){
                String parent = q.poll(); 
                boolean exist = false;
                for(int i=0; i<wordList.size(); i++){
                    String child = wordList.get(i);
                   
                        
                    if(!visited.contains(child) && compareWord(parent, child)){
                        visited.add(child); //hot
                        q.add(child);
                        if(child.equals(endWord)){
                            return tseq+1;
                        }
                    }
                }
                size --;
            }
            //level ended
            tseq++;
            // set.removeAll(); //Emptying Visited set at every level (need to revisit)
        }
        return 0;
    }

    private boolean compareWord(String beginWord, String currWord){ //hit hot emptySet
        int diff = 0;

        if(beginWord.length() != currWord.length() )
            return false;
        
        int len = beginWord.length();
        for(int i=0; i<len; i++ ){
            if( beginWord.charAt(i) != currWord.charAt(i) ){
                diff++;

                if(diff > 1)
                    return false;
            }
        }
        return true;
    }
}
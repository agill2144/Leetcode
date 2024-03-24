class Solution {
    Set<String> visited = new HashSet<>();
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        //start word can not be end word
        if(beginWord == null || beginWord.length() == 0 || endWord == null || endWord.length() == 0) return 0;

        wordList.add(beginWord);
        HashMap<String, List<String>> adjList = new HashMap<>();
        for (int i = 0; i < wordList.size(); i++) {
            String parent = wordList.get(i);
            for (int j = 0; j < wordList.size(); j++) {
                if (i == j) {continue;}
                String child = wordList.get(j);
                boolean isValidChild = compareWord(parent, child);
                if (isValidChild) {
                    List<String> pList = adjList.getOrDefault(parent, new ArrayList<>());
                    pList.add(child);
                    adjList.put(parent, pList);
                }
            }
        }
        Queue<String> q = new LinkedList<>();
        int size = 0;
        int tseq = 1;
        q.add(beginWord); 
        visited.add(beginWord);
        

        while(!q.isEmpty()){ // n*n ; n+n
            size = q.size();
            while(size>0){
                String parent = q.poll(); 
                List<String> childs = adjList.get(parent);
                if (childs == null) {continue;}
                for(String child: childs) {
                    if(!visited.contains(child)){
                        visited.add(child); 
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
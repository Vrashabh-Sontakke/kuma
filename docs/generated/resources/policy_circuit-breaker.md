## CircuitBreaker

- `sources` (required, repeated)

    List of selectors to match dataplanes that are sources of traffic.    
    
    - `match` (optional)
    
        Tags to match, can be used for both source and destinations

- `destinations` (required, repeated)

    List of selectors to match services that are destinations of traffic.    
    
    - `match` (optional)
    
        Tags to match, can be used for both source and destinations

- `conf` (required)    
    
    - `interval` (optional)
    
        Time interval between ejection analysis sweeps    
    
    - `baseejectiontime` (optional)
    
        The base time that a host is ejected for. The real time is equal to the
        base time multiplied by the number of times the host has been ejected    
    
    - `maxejectionpercent` (optional)
    
        The maximum percent of an upstream cluster that can be ejected due to
        outlier detection, has to be in [0 - 100] range    
    
    - `splitexternalandlocalerrors` (optional)
    
        Enables Split Mode in which local and external errors are distinguished    
    
    - `detectors` (optional)    
        
        - `totalerrors` (optional)
        
            Errors with status code 5xx and locally originated errors, in Split
            Mode - just errors with status code 5xx    
            
            - `consecutive` (optional)    
        
        - `gatewayerrors` (optional)
        
            Subset of 'total' related to gateway errors (502, 503 or 504 status
            code)    
            
            - `consecutive` (optional)    
        
        - `localerrors` (optional)
        
            Takes into account only in Split Mode, number of locally originated
            errors    
            
            - `consecutive` (optional)    
        
        - `standarddeviation` (optional)    
            
            - `requestvolume` (optional)
            
                Ignore hosts with less number of requests than 'requestVolume'    
            
            - `minimumhosts` (optional)
            
                Won't count success rate for cluster if number of hosts with required
                'requestVolume' is less than 'minimumHosts'    
            
            - `factor` (optional)
            
                Resulting threshold = mean - (stdev * factor)    
        
        - `failure` (optional)    
            
            - `requestvolume` (optional)
            
                Ignore hosts with less number of requests than 'requestVolume'    
            
            - `minimumhosts` (optional)
            
                Won't count success rate for cluster if number of hosts with required
                'requestVolume' is less than 'minimumHosts'    
            
            - `threshold` (optional)
            
                Eject host if failure percentage of a given host is greater than or
                equal to this value, has to be in [0 - 100] range    
    
    - `thresholds` (optional)    
        
        - `maxconnections` (optional)
        
            The maximum number of connections that Envoy will make to the upstream
            cluster. If not specified, the default is 1024.    
        
        - `maxpendingrequests` (optional)
        
            The maximum number of pending requests that Envoy will allow to the
            upstream cluster. If not specified, the default is 1024.    
        
        - `maxretries` (optional)
        
            The maximum number of parallel retries that Envoy will allow to the
            upstream cluster. If not specified, the default is 3.    
        
        - `maxrequests` (optional)
        
            The maximum number of parallel requests that Envoy will make to the
            upstream cluster. If not specified, the default is 1024.

